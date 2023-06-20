#include <linux/module.h>
#include <linux/platform_device.h>
#include <linux/uio_driver.h>
#include <linux/slab.h>

struct uio_info kpart_info = {
    .name = "kpart",
    .version = "0.1",
    .irq = UIO_IRQ_NONE,
};
static int drv_kpart_probe(struct device *dev);
static int drv_kpart_remove(struct device *dev);
static struct device_driver uio_dummy_driver = {
    .name = "kpart",
    .bus = &platform_bus_type,
    .probe = drv_kpart_probe,
    .remove = drv_kpart_remove,
}

static int drv_kpart_probe(struct device *dev) {
    printf("drv_kpart_probe(%p)\n", dev);
    kpart_info.mem[0].addr = (unsigned long)kmalloc(1024,GFP_KERNEL);

    if (kpart_info.mem[0].addr == 0)
        return -ENOMEM
    kpart_info.mem[0].memtype = UIO_MEM_LOGICAL;
    kpart_info.mem[0].size = 1024;

    if (uio_register_device(dev, &kpart_info))
        return -ENODEV;
    return 0;
}

static int drv_kpart_remove(struct device *dev) {
    uio_unregister_device(&kpart_info);
    return 0;
}

static struct platform_device *uio_dummy_device;
