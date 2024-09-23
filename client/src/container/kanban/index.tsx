"use client";

import React from "react";
import Board from "@asseinfo/react-kanban";
import { useBoardData } from "./hooks/useBoardData";

const KanbanBoard = () => {
  const { title, board } = useBoardData();
  return (
    <Board
      allowRemoveLane
      allowRenameColumn
      allowRemoveCard
      onLaneRemove={console.log}
      onCardRemove={console.log}
      onLaneRename={console.log}
      initialBoard={board}
      allowAddCard={{ on: "top" }}
      onNewCardConfirm={(draftCard: any) => ({
        id: new Date().getTime(),
        ...draftCard,
      })}
      onCardNew={console.log}
    />
  );
};

export default KanbanBoard;
