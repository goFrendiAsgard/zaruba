import sys, traceback

def cli(fn):
    args = []
    kwargs = {}
    arguments = sys.argv[1:]
    for index, argv in enumerate(arguments):
        if argv == '' and ''.join(arguments[index:]) == '':
            break
        pair = argv.split('=')
        if len(pair) < 2:
            args.append(argv)
            continue
        key = pair[0]
        val = '='.join(pair[1:])
        kwargs[key] = val
    # define inner function
    def wrapped_fn():
        try:
            return fn(*args, **kwargs)
        except Exception as e:
            print(e)
            print({"args": args, "kwargs": kwargs})
            traceback.print_exc()
            sys.exit(1)
    return wrapped_fn

