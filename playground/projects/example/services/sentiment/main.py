import numpy as np
from classifiers.classifier import getNewClassifier

labels = ["good", "bad"]
X = np.random.rand(50, 5, 300)  # random word embedding
Y = np.random.rand(50, 2)  # random labels, 0: good, 1: bad

X_train = X[:30]
Y_train = Y[:30]
X_test = X[30:]
Y_test = Y[30:]


if __name__ == "__main__":
    classifier = getNewClassifier("dummy")
    classifier.train(X_train, Y_train, X_test, Y_test)
    # at this point we usually pickle the classifier, so that we can use it later
    prediction = classifier.predict([X_test[0]])
    for index, val in enumerate(prediction[0]):
        print(labels[index], val)
