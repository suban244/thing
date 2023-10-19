# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import graderrequest_pb2 as graderrequest__pb2


class GraderRequestServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GradeFile = channel.unary_unary(
                '/GraderRequestPackage.GraderRequestService/GradeFile',
                request_serializer=graderrequest__pb2.File.SerializeToString,
                response_deserializer=graderrequest__pb2.Status.FromString,
                )


class GraderRequestServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GradeFile(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_GraderRequestServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GradeFile': grpc.unary_unary_rpc_method_handler(
                    servicer.GradeFile,
                    request_deserializer=graderrequest__pb2.File.FromString,
                    response_serializer=graderrequest__pb2.Status.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'GraderRequestPackage.GraderRequestService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class GraderRequestService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GradeFile(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/GraderRequestPackage.GraderRequestService/GradeFile',
            graderrequest__pb2.File.SerializeToString,
            graderrequest__pb2.Status.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)