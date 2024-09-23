import { BoardEntity } from "@/modules/board/entities/board-entities";
import { useMemo } from "react";
export const useBoardData = (defaultBoardData: BoardEntity | undefined) => {
  const board = defaultBoardData;
  const title = board?.name ?? "";

  console.log("board?.columns", board?.columns);
  const transformColumnsData = useMemo(() => {
    const boardResult = (board?.columns || []).map((column) => ({
      id: column.id,
      title: column.name,
      backgroundColor: "#fff",
      cards: (column?.tasks || []).map((task) => ({
        id: task.id,
        title: task.title,
        description: task.description,
      })),
    }));
    return boardResult;
  }, [defaultBoardData]);

  return {
    title,
    board: { columns: transformColumnsData },
  };
};
