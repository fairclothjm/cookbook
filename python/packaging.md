# python packaging

See [python-packaging](https://python-packaging.readthedocs.io/en/latest/index.html)

### choose a name
- all lowercase
- underscore-separated or no word separators (no hyphens)

### structure
```
mypackage/
    mypackage/
        __init__.py
        hello.py
    setup.py
```

The top level dir is the root of the git repo. The sub dir, also called `mypackage` is the actual python module.

In `__init__.py`:

```python
from .hello import hello_world
```

In `hello.py`:

```python
def hello_world():
    return "hello world"
```


```python
from setuptools import setup

setup(name='mypackage',
      version='0.1',
      description='The mypackage',
      url='http://github.com/test/mypackage',
      author='Flying Circus',
      author_email='flyingcircus@example.com',
      license='MIT',
      packages=['mypackage'],
      zip_safe=False)
```

### install
Install the package with a symlink:
```
pip install -e .
```

Anywhere on the system, this should work:

```python
import mypackage
print(mypackage.hello_world())
```
