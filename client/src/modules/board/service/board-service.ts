import BaseService from "@/config/service/base-service";
import fetcher from "@/config/service/fetcher-service";
import BaseResponse from "@/config/entities/core-entities";
import {
  BoardEntity,
  BoardUpdateEntity,
  CreateTaskRequest,
} from "../entities/board-entities";

class BoardService implements BaseService {
  readonly QUERY_KEY = {
    BOARD: "/board/1",
    CREATE_TASK: "/board/create-task",
    DELETE_TASK: "/board/delete-task",
  };

  fetcher = fetcher.coreFetcher;

  board = () => {
    return this.fetcher.get<BoardEntity>(this.QUERY_KEY.BOARD, {});
  };

  updateBoard = (body: BoardUpdateEntity) => {
    return this.fetcher.put<BaseResponse<any>>(this.QUERY_KEY.BOARD, body);
  };

  createStack = (body: CreateTaskRequest) => {
    return this.fetcher.post<BaseResponse<any>>(
      this.QUERY_KEY.CREATE_TASK,
      body
    );
  };

  deleteTask = (id: number) => {
    return this.fetcher.delete<BaseResponse<any>>(
      this.QUERY_KEY.DELETE_TASK + "/" + id
    );
  };
}

const boardService = new BoardService();

export default boardService;
