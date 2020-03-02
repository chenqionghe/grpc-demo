from concurrent import futures
import logging
import grpc

# 支持新的包
import sys
sys.path.append("lightweight")
import lightweight.gym_pb2_grpc as gym_pb2_grpc
import lightweight.gym_pb2 as gym_pb2


class Gym(gym_pb2_grpc.GymServicer):

    def BodyBuilding(self, request, context):
        print(f"{request.name}在健身, 动作: {list(request.actions)}")
        return gym_pb2.Reply(code=0, msg='ok')


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    gym_pb2_grpc.add_GymServicer_to_server(Gym(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
