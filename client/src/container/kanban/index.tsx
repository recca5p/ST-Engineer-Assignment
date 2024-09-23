"use client";

import React from "react";
import Board from "@asseinfo/react-kanban";
import { useBoardData } from "./hooks/useBoardData";
import { useGetBoard } from "@/modules/board";
import { useCreateTask } from "@/modules/board/application/use-create-task";
import { useQueryClient } from "react-query";
import boardService from "@/modules/board/service/board-service";
import { LoaderCircle } from "lucide-react";
import { useDeleteTask } from "@/modules/board/application/use-delete-task";

const KanbanBoard = () => {
    const ql = useQueryClient();

    const { data, isLoading: isLoadingGetBoard } = useGetBoard();
    const { mutate: handleCreateTask } = useCreateTask();
    const { mutate: handleDeleteTask } = useDeleteTask();

    // Function to handle task creation with dynamic card info
    const funcCallCreateTask = (newCard, columnId) => {
        handleCreateTask(
            {
                columnid: columnId,
                name: newCard.name,
                position: newCard.position || 1,
                description: newCard.description || "No description",
            },
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

    // Function to handle task deletion with dynamic card info
    const funcCallDeleteTask = (cardId) => {
        handleDeleteTask(
            cardId,
            {
                onSuccess: (res) => {
                    ql.invalidateQueries([boardService.QUERY_KEY.BOARD]);
                    console.log("Task deleted:", res);
                },
                onError: (error) => {
                    console.log("Error deleting task:", error);
                },
            }
        );
    };

    const { title, board } = useBoardData(data);

/*
    console.table("board", board);
*/

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
                        allowRemoveLane
                        allowRenameColumn
                        allowRemoveCard
                        onLaneRemove={console.log}
                        onCardRemove={(card) => funcCallDeleteTask(card.id)}
                        onLaneRename={console.log}
                        onCardDragEnd={console.log}
                        onColumnDragEnd={console.log}
                        initialBoard={board}
                        allowAddCard={{ on: "top" }}
                        onNewCardConfirm={(draftCard) => ({
                            id: new Date().getTime(),
                            title: draftCard.title,
                            description: draftCard.description || "No description",
                        })}
                        onCardNew={console.log}
                    />
                </>
            )}
        </div>
    );
};

export default KanbanBoard;
