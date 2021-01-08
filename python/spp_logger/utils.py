import immutables


def context_to_dict(context: immutables.Map) -> dict:
    context_dict = {}
    for key in context.keys():
        context_dict[key] = context.get(key)
    return context_dict


def dict_to_context(context_dict: dict) -> immutables.Map:
    return immutables.Map(**context_dict)
