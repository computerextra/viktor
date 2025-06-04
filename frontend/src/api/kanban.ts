import {
  CreatePost as APICreatePost,
  DeleteBoard as APIDeleteBoard,
  DeletePost as APIDeletePost,
  UpdatePost as APIUpdatePost,
  CreateKanban,
  GetKanbanBoardsFromUser,
  GetKanbanBord,
  UpdateKanban,
} from "@bindings/viktor/backend/app";
import type { Kanban } from "@bindings/viktor/db/models";
import { z } from "zod";
import { CheckSession } from "./userdata";

export const CreateBoard = async (BoardName: string): Promise<boolean> => {
  const session = await CheckSession();
  if (session == null) return false;
  if (session.Id == null) return false;
  if (session.UserId == null) return false;
  return CreateKanban(session.UserId, BoardName);
};

export const PostParams = z.object({
  Name: z.string(),
  Description: z.string().optional(),
  Importance: z.string(),
  Status: z.string(),
});

type PostParams = z.infer<typeof PostParams>;

export const CreatePost = async (
  BoardId: string,
  Params: PostParams
): Promise<boolean> => {
  return await APICreatePost(
    Params.Name,
    Params.Status,
    Params.Importance,
    Params.Description ? Params.Description : null,
    BoardId
  );
};

export const GetBoardsFromUser = async (): Promise<Array<Kanban> | null> => {
  const session = await CheckSession();
  if (session == null) return null;
  if (session.Id == null) return null;
  if (session.UserId == null) return null;
  return await GetKanbanBoardsFromUser(session.UserId);
};

export const GetBoard = async (BoardId: string): Promise<Kanban | null> => {
  return await GetKanbanBord(BoardId);
};

export const UpdateBoard = async (
  BoardId: string,
  Name: string
): Promise<boolean> => {
  return await UpdateKanban(BoardId, Name);
};

export const UpdatePost = async (
  PostId: string,
  Params: PostParams
): Promise<boolean> => {
  return await APIUpdatePost(
    PostId,
    Params.Name,
    Params.Status,
    Params.Importance,
    Params.Description ? Params.Description : null
  );
};

export const DeletePost = async (PostId: string): Promise<boolean> => {
  return await APIDeletePost(PostId);
};

export const DeleteBoard = async (BoardId: string): Promise<boolean> => {
  return await APIDeleteBoard(BoardId);
};
