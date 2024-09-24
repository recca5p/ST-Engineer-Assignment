import { FetcherError } from "@/config/domain/fetcher";
import BaseResponse from "@/config/entities/core-entities";
import { useMutation } from "react-query";
import { BoardUpdateEntity } from "../entities/board-entities";
import boardService from "../service/board-service";

export const useUpdateBoard = () =>
  useMutation<BaseResponse<any>, FetcherError, BoardUpdateEntity>({
    mutationKey: [boardService.QUERY_KEY.BOARD],
    mutationFn: async (variables) => {
      const result = await boardService.updateBoard(variables);

      return result.data;
    },
  });
