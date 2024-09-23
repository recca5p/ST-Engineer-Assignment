import { useQuery } from "react-query";

import { BoardEntity } from "../entities/board-entities";
import boardService from "../service/board-service";

export const useGetBoard = () =>
  useQuery<BoardEntity>({
    queryKey: [boardService.QUERY_KEY.BOARD],
    queryFn: async () => {
      const result = await boardService.board();

      return result.data;
    },
    refetchOnWindowFocus: true,
  });
