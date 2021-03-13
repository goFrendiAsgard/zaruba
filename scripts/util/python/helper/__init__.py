import sys, traceback

def cli(fn):
    args = []
    kwargs = {}
    for index, argv in enumerate(sys.argv[1:]):
        if argv == '' and ''.join(argv[index:]) == '':
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
            traceback.print_exc()
            sys.exit(1)
    return wrapped_fn

