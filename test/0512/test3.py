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
        # model = models.vgg16(False, False)
        model = models.mobilenet_v2(weights=None, progress=False)
        mem('After loading model')

    # load state_dict
    with timer.Timer(fotmat='  -> Elapsed time : %.5f ms'):
        # state_dict = torch.load('../vgg16.pth')
        state_dict = torch.load('../mobilenetv2.pth')
        mem('After loading state_dict')

    # set state_dict to model
    with timer.Timer(fotmat='  -> Elapsed time : %.5f ms'):
        model.load_state_dict(state_dict)
        mem('After connecting model and state_dict')

    with timer.Timer(fotmat='  -> Elapsed time : %.5f ms'):
        del state_dict
        mem('After releasing state_dict')

    input()

    # predict
    with timer.Timer(fotmat='  -> Elapsed time : %.5f ms'):
        input_img = Image.open('/Users/freckie/prj/shmfaas/test/0623/fn-shmfaas/dog-224.jpg').convert('RGB')
        input_tensor = torch.unsqueeze(ToTensor()(np.array(input_img)), 0)
        model.eval()
        result = model(input_tensor)
        print(torch.argmax(result, dim=1))