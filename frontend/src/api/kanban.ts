import {
  CreateBoard as APICreateBoard,
  CreatePost as APICreatePost,
  DeleteBoard as APIDeleteBoard,
  DeletePost as APIDeletePost,
  GetBoard as APIGetBoard,
  GetBoardFromUser as APIGetBoardFromUser,
  UpdateBoard as APIUpdateBoard,
  UpdatePost as APIUpdatePost,
} from "@wails/go/main/App";
import type { db } from "@wails/go/models";
import { z } from "zod";
import { CheckSession } from "./userdata";

export const CreateBoard = async (BoardName: string): Promise<boolean> => {
  const session = await CheckSession();
  if (session == null) return false;
  if (session.Id == null) return false;
  if (session.UserId == null) return false;
  return APICreateBoard(session.UserId, BoardName);
};

export const PostParams = z.object({
  Name: z.string(),
  Description: z.string().optional(),
  Importance: z.string(),
  Status: z.string(),
});

type PostParams = z.infer<typeof PostParams>;

export const CreatePost = async (
  BoardId: number,
  Params: PostParams
): Promise<boolean> => {
  return await APICreatePost(
    BoardId,
    Params.Name,
    Params.Description,
    Params.Importance,
    Params.Status
  );
};

export const GetBoardsFromUser = async (): Promise<Array<db.Kanban> | null> => {
  const session = await CheckSession();
  if (session == null) return null;
  if (session.Id == null) return null;
  if (session.UserId == null) return null;
  return await APIGetBoardFromUser(session.UserId);
};

export const GetBoard = async (BoardId: number): Promise<db.Kanban> => {
  return await APIGetBoard(BoardId);
};

export const UpdateBoard = async (
  BoardId: number,
  Name: string
): Promise<boolean> => {
  return await APIUpdateBoard(BoardId, Name);
};

export const UpdatePost = async (
  PostId: number,
  Params: PostParams
): Promise<boolean> => {
  return await APIUpdatePost(
    PostId,
    Params.Name,
    Params.Description,
    Params.Status,
    Params.Importance
  );
};

export const DeletePost = async (PostId: number): Promise<boolean> => {
  return await APIDeletePost(PostId);
};

export const DeleteBoard = async (BoardId: number): Promise<boolean> => {
  return await APIDeleteBoard(BoardId);
};
