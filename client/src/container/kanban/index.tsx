"use client";

import React from "react";
import Board from "@asseinfo/react-kanban";
import { useBoardData } from "./hooks/useBoardData";
import { useGetBoard, useUpdateBoard } from "@/modules/board";
import { useCreateTask } from "@/modules/board/application/use-create-task";
import { useQueryClient } from "react-query";
import boardService from "@/modules/board/service/board-service";
import { LoaderCircle } from "lucide-react";

const KanbanBoard = () => {
  const ql = useQueryClient();

  const { data, isLoading: isLoadingGetBoard } = useGetBoard();
  const { mutate: handleUpdateBoard, isLoading: isUpdating } = useUpdateBoard();

  // Function to handle task creation with dynamic card info
  const funcCallCreateTask = (newCard, columnId: number) => {
    console.log("newCard", newCard, "columnId", columnId);
    const defaultCards = (data?.columns || []).filter(
      (column) => column.id === columnId
    )[0].tasks;
    const cards = newCard.columns.filter((column) => column.id === columnId)[0]
      .cards;

    const transformCards = (cards ?? []).map((card) => ({
      ...((defaultCards ?? []).map((c) => c.id).includes(card.id)
          ? { id: card.id }
          : {}),
      title: card.title,
      description: card.description,
    }));



    handleUpdateBoard(
      newCard.columns.map((column) => ({
        id: column.id,
        title: column.title,
        cards: column.id === columnId ? transformCards : column.cards,
      })),
      {
        onSuccess: (res) => {
          ql.invalidateQueries([boardService.QUERY_KEY.BOARD]);
          console.log("Task created:", res);
        },
        onError: (error) => {
          console.log("Error creating task:", error);
        },
      }
    );
  };
  const funcCallUpdateBoard = (card) => {
    handleUpdateBoard(card.columns, {
      onSuccess: (res) => {
        ql.invalidateQueries([boardService.QUERY_KEY.BOARD]);
        console.log("Task created:", res);
      },
      onError: (error) => {
        console.log("Error creating task:", error);
      },
    });
  };


  const { title, board } = useBoardData(data);

  return (
    <div className="flex flex-col justify-center h-full items-center p-5">
      {isLoadingGetBoard ? (
        <div className="h-full w-full">
          <LoaderCircle className="animate-spin" />{" "}
        </div>
      ) : (
        <>
          <div className="text-white text-[28px] font-bold">{title}</div>
          <Board
            allowRemoveLane={!(isLoadingGetBoard || isUpdating)}
            allowRenameColumn={!(isLoadingGetBoard || isUpdating)}
            allowRemoveCard={!(isLoadingGetBoard || isUpdating)}
            onLaneRemove={console.log}
            onCardRemove={(card) => funcCallUpdateBoard(card)}
            onLaneRename={(card) => funcCallUpdateBoard(card)}
            onColumnRename={(card) => funcCallUpdateBoard(card)}
            onColumnDragEnd={(card) => funcCallUpdateBoard(card)}
            onCardDragEnd={(card) => funcCallUpdateBoard(card)}
            initialBoard={board}
            allowAddCard={{ on: "top" }}
            disableColumnDrag={isLoadingGetBoard || isUpdating}
            disableCardDrag={isLoadingGetBoard || isUpdating}
            onNewCardConfirm={(draftCard) => ({
              title: draftCard.title,
              description: draftCard.description || "No description",
            })}
            onCardNew={(newCard, column) =>
              funcCallCreateTask(newCard, column.id)
            } // Pass new card and column info
          />
        </>
      )}
    </div>
  );
};

export default KanbanBoard;
