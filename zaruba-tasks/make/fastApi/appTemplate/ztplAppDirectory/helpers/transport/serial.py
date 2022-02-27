from typing import Any, Callable
import jsons

def create_json_string_encoder(encoding_type: str = 'utf-8') -> Callable[[Any], bytes]:
    def encode(data: Any) -> bytes:
        serialized = jsons.dumps(data).encode(encoding_type)
        if type(serialized) == bytes:
            return serialized.decode(encoding_type)
        return serialized
    return encode

def create_json_byte_encoder(encoding_type: str = 'utf-8') -> Callable[[Any], bytes]:
    def encode(data: Any) -> bytes:
        serialized = jsons.dumps(data).encode(encoding_type)
        if type(serialized) == bytes:
            return serialized
        return bytes(serialized, encoding_type)
    return encode

def create_json_decoder(encoding_type: str = 'utf-8') -> Callable[[bytes], Any]:
    def decode(serialized: bytes) -> Any:
        if type(serialized) == bytes:
            return jsons.loads(serialized.decode(encoding_type))
        return jsons.loads(serialized)
    return decode