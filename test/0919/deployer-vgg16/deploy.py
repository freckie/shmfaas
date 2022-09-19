import os
import json
import pickle

from flask import Flask

import numpy as np
import torch
import torchvision.models as models
from torchvision.transforms import ToTensor
from PIL import Image

import shmtorch

if __name__ == "__main__":
    env = os.environ
    addr = env['NODE_NAME'] + ':20000'
    model_name = env['SHMM_NAME']
    tag_name = env['TAG_NAME']
    shmname = env['SHMNAME']

    model = models.vgg16(True, True)
    shm, metadata = shmtorch.x_save_states(model, shmname)
    shmtorch.x_apply_to_shmm(addr, model_name, tag_name, metadata)

    print('Finished')
