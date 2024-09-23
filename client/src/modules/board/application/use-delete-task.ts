import { FetcherError } from "@/config/domain/fetcher";
import BaseResponse from "@/config/entities/core-entities";
import { useMutation } from "react-query";
import { CreateTaskRequest } from "../entities/board-entities";
import boardService from "../service/board-service";

export const useDeleteTask = () =>
    useMutation<BaseResponse<any>, FetcherError, number>({
        mutationKey: [boardService.QUERY_KEY.DELETE_TASK],
        mutationFn: async (id) => {
            const result = await boardService.deleteTask(id);

            return result.data;
        },
    });
