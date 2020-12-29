from fastapi import FastAPI, HTTPException
import transport, schema

def init(app: FastAPI, mb: transport.MessageBus):
    print('Init {} route handlers'.format('module'))
