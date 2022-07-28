import time
lib_start = time.perf_counter()

import os, psutil

proc = psutil.Process(os.getpid())
curr = 0
msgfmt = '$ %10d | %s'

def mem(msg=''):
    print(msgfmt % (proc.memory_info().rss, msg))

import torch
import torchvision.models as models
from torchvision.transforms import ToTensor

import numpy as np
from PIL import Image

import pickle
from multiprocessing.shared_memory import SharedMemory

import shmtorch
import timer

if __name__ == '__main__':
    mem('Initialized')
    print('  -> Elapsed time : %.5f ms' % (time.perf_counter() - lib_start))

    # model load
    with timer.Timer(fotmat='  -> Elapsed time : %.5f ms'):
        # model_skeleton = models.vgg16(False, False)
        model_skeleton = models.mobilenet_v2(weights=None, progress=False)
        mem('After loading model')

    # load metadata from pickle
    with timer.Timer(fotmat='  -> Elapsed time : %.5f ms'):
        metadata: shmtorch.XMetadata
        # with open('/Users/freckie/prj/shmfaas/test/0512/vgg16-meta', 'rb') as f:
        with open('./mobilenetv2-meta', 'rb') as f:
            metadata = pickle.load(f)
        mem('After loading the metadata')

    # load tensors into the model
    with timer.Timer(fotmat='  -> Elapsed time : %.5f ms'):
        shm, model = shmtorch.x_load_states(model_skeleton, metadata)
        model.eval()
        mem('After loading tensors')

    # del the skeleton model
    with timer.Timer(fotmat='  -> Elapsed time : %.5f ms'):
        del model_skeleton
        del metadata
        mem('After releasing the skeleton model')

    input()
    mem('After input')

    # predict
    with timer.Timer(fotmat='  -> Elapsed time : %.5f ms'):
        input_img = Image.open('/Users/freckie/prj/shmfaas/test/0623/fn-shmfaas/dog-224.jpg').convert('RGB')
        input_tensor = torch.unsqueeze(ToTensor()(np.array(input_img)), 0)
        model.eval()
        result = model(input_tensor)
        print(torch.argmax(result, dim=1))

    input()

    shm.close()
        