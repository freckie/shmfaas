from time import perf_counter

class Timer:
    def __init__(self, fotmat, f=print):
        self.t = None
        self.f = f
        self.fmt = fotmat
    
    def __enter__(self):
        self.t = perf_counter()

    def __exit__(self, type, value, traceback):
        self.f(self.fmt % (perf_counter()-self.t))