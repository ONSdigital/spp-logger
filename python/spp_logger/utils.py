from typing import Optional

import immutables


def context_to_dict(context: immutables.Map) -> dict:
    keys = [key for key in context.keys()]
    context_dict = {}
    for key in sorted(keys):
        context_dict[key] = context.get(key)
    return context_dict


def dict_to_context(context_dict: Optional[dict]) -> Optional[immutables.Map]:
    if not context_dict:
        return None
    return immutables.Map(**context_dict)
