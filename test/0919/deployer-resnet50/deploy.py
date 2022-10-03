import os

import numpy as np
import torchvision.models as models

import shmtorch

if __name__ == "__main__":
    env = os.environ
    addr = env['NODE_NAME'] + ':20000'
    model_name = env['SHMM_NAME']
    tag_name = env['TAG_NAME']

    model = models.resnet50(True, True)
    model.eval()

    shmsize = shmtorch.x_calc_bytes(model)
    shmname = shmtorch.x_create_shm(addr, model_name, tag_name, shmsize)
    shm, metadata = shmtorch.x_save_states(model, shmname)
    shmtorch.x_apply_shm(addr, model_name, tag_name, metadata)
    shm.close()

    print('Finished')
