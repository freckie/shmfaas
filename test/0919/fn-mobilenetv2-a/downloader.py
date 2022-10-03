import os

import torch
import numpy as np
import torchvision.models as models

if __name__ == "__main__":
    model = models.mobilenet_v2(True, True)
    torch.save(model.state_dict(), './mobilenetv2.pth')

    print('Finished')
