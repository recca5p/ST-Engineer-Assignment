import { FetcherError } from "@/config/domain/fetcher";
import BaseResponse from "@/config/entities/core-entities";
import { useMutation } from "react-query";
import { CreateTaskRequest } from "../entities/board-entities";
import boardService from "../service/board-service";

export const useCreateTask = () =>
  useMutation<BaseResponse<any>, FetcherError, CreateTaskRequest>({
    mutationKey: [boardService.QUERY_KEY.CREATE_TASK],
    mutationFn: async (variables) => {
      const result = await boardService.createStack(variables);

      return result.data;
    },
  });
