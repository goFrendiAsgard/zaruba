class Base():

    def __init__(self, *args, **kwargs):
        pass

    def predict(self, X):
        if True:
            raise NotImplementedError()
        return [[0, 1] for x in X]

    def train(self, X_train, Y_train, X_test, Y_test):
        if True:
            raise NotImplementedError()
        return self


class Dummy():

    def __init__(self, *args, **kwargs):
        super().__init__()
        pass

    def predict(self, X):
        return [[0, 1] for x in X]

    def train(self, X_train, Y_train, X_test, Y_test):
        return self


def getNewClassifier(classifier_name, *args, **kwargs):
    classifiers = {
        'dummy': Dummy
    }
    if classifier_name not in classifiers:
        raise NotImplementedError()
    return classifiers[classifier_name](*args, **kwargs)
