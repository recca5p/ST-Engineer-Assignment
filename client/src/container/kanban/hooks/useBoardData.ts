import rawData from "@/container/kanban/data/mocking.json";
import { useMemo } from "react";
export const useBoardData = () => {
  const board = rawData.boards[0];
  const title = board.name;

  const transformColumnsData = useMemo(() => {
    const boardResult = board.columns.map((column) => ({
      id: column.id,
      title: column.name,
      backgroundColor: "#fff",
      cards: column.tasks.map((task) => ({
        id: task.id,
        title: task.title,
      })),
    }));
    return boardResult;
  }, []);

  return {
    title,
    board: { columns: transformColumnsData },
  };
};
