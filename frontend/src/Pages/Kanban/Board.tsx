import { CreatePost, DeletePost, GetBoard, UpdatePost } from "@/api/kanban";
import BackButton from "@/components/BackButton";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Textarea } from "@/components/ui/textarea";
import { cn } from "@/lib/utils";
import type { db } from "@wails/go/models";
import { GripVertical, Plus, X } from "lucide-react";
import type React from "react";
import { useEffect, useState } from "react";
import { useParams } from "react-router";

interface Column {
  id: string;
  title: string;
  posts: db.Post[];
}

export default function KanbanBoard() {
  const { id } = useParams();

  const [board, setBoard] = useState<db.Kanban | undefined>(undefined);
  const [columns, setColumns] = useState<Column[] | undefined>(undefined);

  const [draggedTask, setDraggedTask] = useState<db.Post | null>(null);
  const [draggedFrom, setDraggedFrom] = useState<string | null>(null);
  const [newTaskTitle, setNewTaskTitle] = useState("");
  const [showAddTask, setShowAddTask] = useState<string | null>(null);

  useEffect(() => {
    (async () => {
      if (id == null) return;
      const res = await GetBoard(parseInt(id));
      setBoard(res);
      const todo: db.Post[] = [];
      const inprogress: db.Post[] = [];
      const review: db.Post[] = [];
      const done: db.Post[] = [];

      res.Posts.forEach((x) => {
        switch (x.Status) {
          case "todo":
            todo.push(x);
            break;
          case "inprogress":
            inprogress.push(x);
            break;
          case "review":
            review.push(x);
            break;
          case "done":
            done.push(x);
            break;
        }
      });

      setColumns([
        {
          id: "todo",
          title: "To Do",
          posts: todo,
        },
        {
          id: "inprogress",
          title: "In Arbeit",
          posts: inprogress,
        },
        {
          id: "review",
          title: "Prüfen",
          posts: review,
        },
        {
          id: "done",
          title: "Fertig",
          posts: done,
        },
      ]);
    })();
  }, [id]);

  const handleDragStart = (
    e: React.DragEvent,
    task: db.Post,
    columnId: string
  ) => {
    setDraggedTask(task);
    setDraggedFrom(columnId);
    e.dataTransfer.effectAllowed = "move";
  };

  const handleDragOver = (e: React.DragEvent) => {
    e.preventDefault();
    e.dataTransfer.dropEffect = "move";
  };

  const handleDrop = async (e: React.DragEvent, targetColumnId: string) => {
    e.preventDefault();

    if (!draggedTask || !draggedFrom) return;

    if (draggedFrom === targetColumnId) {
      setDraggedTask(null);
      setDraggedFrom(null);
      return;
    }

    // Update DB
    const res = await UpdatePost(draggedTask.ID, {
      Importance: draggedTask.Importance,
      Name: draggedTask.Name,
      Description: draggedTask.Description,
      Status: targetColumnId,
    });
    if (res) {
      location.reload();
    } else {
      alert("Komischer Fehler!");
    }
  };

  const addTask = async (status: string) => {
    if (!newTaskTitle.trim()) return;
    if (board == null) return;
    if (board.ID == null) return;
    const res = await CreatePost(board.ID, {
      Importance: "medium",
      Name: newTaskTitle,
      Description: "",
      Status: status,
    });

    if (res) {
      location.reload();
    } else {
      alert("Fehler beim Speichern");
    }
  };

  const deleteTask = async (taskId: number) => {
    const res = await DeletePost(taskId);

    if (res) {
      location.reload();
    } else {
      alert("Fehler beim löschen");
    }
  };

  return (
    <div className="">
      <BackButton href="/Kanban" />
      <div className="max-w-7xl mx-auto">
        <h1 className="text-3xl font-bold text-gray-900 mb-8 text-center">
          {board?.Name}
        </h1>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          {columns?.map((column) => (
            <div
              key={column.id}
              className="bg-gray-100 rounded-lg p-4 min-h-[600px]"
              onDragOver={handleDragOver}
              onDrop={async (e) => await handleDrop(e, column.id)}
            >
              <div className="flex items-center justify-between mb-4">
                <h2 className="font-semibold text-gray-800">{column.title}</h2>
                <Badge variant="secondary" className="text-xs">
                  {column.posts.length}
                </Badge>
              </div>

              <div className="space-y-3">
                {column.posts.map((task) => (
                  <Card
                    key={task.ID}
                    className="cursor-move hover:shadow-md transition-shadow bg-white"
                    draggable
                    onDragStart={(e) => handleDragStart(e, task, column.id)}
                  >
                    <CardHeader className="pb-2">
                      <div className="flex items-start justify-between">
                        <div className="flex items-center gap-2">
                          <GripVertical className="h-4 w-4 text-gray-400" />
                          <PostEditForm post={task} />
                        </div>
                        <Button
                          variant="ghost"
                          size="sm"
                          className="h-6 w-6 p-0 hover:bg-red-100"
                          onClick={async () => await deleteTask(task.ID)}
                        >
                          <X className="h-3 w-3" />
                        </Button>
                      </div>
                    </CardHeader>
                    <CardContent className="pt-0">
                      {task.Description && (
                        <p className="text-xs text-gray-600 mb-3">
                          {task.Description}
                        </p>
                      )}
                      <div className="flex items-center justify-between">
                        <Badge
                          className={cn(
                            "text-xs",
                            task.Importance == "low"
                              ? "bg-green-100 text-green-800"
                              : task.Importance == "medium"
                              ? "bg-yellow-100 text-yellow-800"
                              : "bg-red-100 text-red-800"
                          )}
                        >
                          {task.Importance}
                        </Badge>
                      </div>
                    </CardContent>
                  </Card>
                ))}
              </div>

              {showAddTask === column.id ? (
                <div className="mt-3 space-y-2">
                  <Input
                    placeholder="Enter task title..."
                    value={newTaskTitle}
                    onChange={(e) => setNewTaskTitle(e.target.value)}
                    onKeyPress={async (e) =>
                      e.key === "Enter" && (await addTask(column.id))
                    }
                    autoFocus
                  />
                  <div className="flex gap-2">
                    <Button
                      size="sm"
                      onClick={async () => await addTask(column.id)}
                      className="flex-1"
                    >
                      Add Task
                    </Button>
                    <Button
                      size="sm"
                      variant="outline"
                      onClick={() => {
                        setShowAddTask(null);
                        setNewTaskTitle("");
                      }}
                    >
                      Cancel
                    </Button>
                  </div>
                </div>
              ) : (
                <Button
                  variant="ghost"
                  className="w-full mt-3 border-2 border-dashed border-gray-300 hover:border-gray-400"
                  onClick={() => setShowAddTask(column.id)}
                >
                  <Plus className="h-4 w-4 mr-2" />
                  Add Task
                </Button>
              )}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

function PostEditForm({ post }: { post: db.Post }) {
  const [name, setName] = useState<string | undefined>(undefined);
  const [desc, setDesc] = useState<string | undefined>(undefined);
  const [prio, setPrio] = useState<"low" | "medium" | "high" | undefined>(
    undefined
  );

  const handleSave = async () => {
    let n: string, p: string, d: string | undefined;
    if (name == null) {
      n = post.Name;
    } else {
      n = name;
    }
    if (prio == null) {
      p = post.Importance;
    } else {
      p = prio;
    }
    if (desc == null) {
      d = post.Description;
    } else {
      d = desc;
    }

    const res = await UpdatePost(post.ID, {
      Importance: p,
      Name: n,
      Description: d,
      Status: post.Status,
    });

    if (res) {
      location.reload();
    } else {
      alert("Fehler beim Speichern");
    }
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <h3 className="font-medium text-sm hover:underline">{post.Name}</h3>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>"{post.Name}" bearbeiten</DialogTitle>
        </DialogHeader>
        <div className="grid gap-4 py-4">
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="name" className="text-right">
              Name
            </Label>
            <Input
              id="name"
              required
              defaultValue={post.Name}
              onChange={(e) => setName(e.target.value)}
              className="col-span-3"
            />
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="name" className="text-right">
              Priorität
            </Label>
            <Select
              required
              defaultValue={post.Importance}
              onValueChange={(e) => {
                if (e === "low" || e === "medium" || e === "high") setPrio(e);
              }}
            >
              <SelectTrigger className="w-[180px]">
                <SelectValue placeholder="Auswählen" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <SelectItem value="low">Niedrig</SelectItem>
                  <SelectItem value="medium">Mittel</SelectItem>
                  <SelectItem value="high">Hoch</SelectItem>
                </SelectGroup>
              </SelectContent>
            </Select>
          </div>
          <div className="grid grid-cols-1 items-center gap-4">
            <Label htmlFor="username" className="text-right">
              Beschreibung
            </Label>
            <Textarea
              defaultValue={post.Description}
              onChange={(e) => setDesc(e.target.value)}
            />
          </div>
        </div>
        <DialogFooter>
          <Button type="submit" onClick={handleSave}>
            Änderungen speichern
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
