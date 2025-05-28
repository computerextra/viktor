import {
  CreateBoard,
  DeleteBoard,
  GetBoardsFromUser,
  UpdateBoard,
} from "@/api/kanban";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from "@/components/ui/card";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import useSession from "@/hooks/useSession";
import { db } from "@wails/go/models";
import { AlertCircle } from "lucide-react";
import { useEffect, useState } from "react";
import { NavLink } from "react-router";

export default function KanbanBoards() {
  const session = useSession();
  const [boards, setBoards] = useState<Array<db.Kanban> | null>(null);
  const [name, setName] = useState<string | undefined>(undefined);

  useEffect(() => {
    (async () => {
      const res = await GetBoardsFromUser();
      setBoards(res);
    })();
  }, []);

  const handleCreate = async () => {
    if (name == null) return;

    const res = await CreateBoard(name);
    if (res) {
      location.reload();
    } else {
      alert("Fehler beim Speichern");
    }
  };

  const handleRename = async (id: number) => {
    if (name == null) return;

    const res = await UpdateBoard(id, name);
    if (res) {
      location.reload();
    } else {
      alert("Fehler beim Speichern");
    }
  };

  const handleDelete = async (id: number) => {
    const res = await DeleteBoard(id);
    if (res) {
      location.reload();
    } else {
      alert("Fehler beim Löschen");
    }
  };

  return (
    <>
      <h1 className="text-center">Deine Kanban Boards</h1>
      {!session || session.UserId == null ? (
        <Alert variant={"destructive"} className="my-5">
          <AlertCircle className="h-4 w-4" />
          <AlertTitle>Nicht angemeldet</AlertTitle>
          <AlertDescription>
            Um dieses Feature zu nutzen musst du angemeldet sein!
            <Button asChild>
              <NavLink to="/Anmelden">Hier anmelden</NavLink>
            </Button>
          </AlertDescription>
        </Alert>
      ) : (
        <Dialog>
          <DialogTrigger asChild>
            <Button>Neues Board anlegen</Button>
          </DialogTrigger>
          <DialogContent className="sm:max-w-[425px]">
            <DialogHeader>
              <DialogTitle>Neues Board anlegen</DialogTitle>
              <DialogDescription>
                Lege hier ein neues Board an.
              </DialogDescription>
            </DialogHeader>
            <div className="grid gap-4 py-4">
              <div className="grid grid-cols-4 items-center gap-4">
                <Label htmlFor="name" className="text-right">
                  Name
                </Label>
                <Input
                  id="name"
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                  className="col-span-3"
                />
              </div>
            </div>
            <DialogFooter>
              <Button type="submit" onClick={handleCreate}>
                Anlegen
              </Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      )}
      <div className="mt-12 grid grid-cols-2 gap-4">
        {boards?.map((x) => (
          <Card key={x.ID}>
            <CardHeader>{x.Name}</CardHeader>
            <CardContent>
              <div className="grid grid-cols-2 gap-2 mb-3 max-w-[50%]">
                <span>Erstellt</span>
                <span className="text-end">
                  {new Date(x.CreatedAt).toLocaleDateString("de-DE", {
                    day: "2-digit",
                    month: "2-digit",
                    year: "2-digit",
                    hour: "2-digit",
                    minute: "2-digit",
                  })}
                </span>
                <span>Letztes Update</span>
                <span className="text-end">
                  {new Date(x.UpdatedAt).toLocaleDateString("de-DE", {
                    day: "2-digit",
                    month: "2-digit",
                    year: "2-digit",
                    hour: "2-digit",
                    minute: "2-digit",
                  })}
                </span>
              </div>
              <Button asChild>
                <NavLink to={"/Kanban/" + x.ID}>Öffnen</NavLink>
              </Button>
            </CardContent>
            <CardFooter>
              <div className="flex flex-row justify-around w-full">
                <Dialog>
                  <DialogTrigger asChild>
                    <Button variant={"secondary"}>Namen bearbeiten</Button>
                  </DialogTrigger>
                  <DialogContent className="sm:max-w-[425px]">
                    <DialogHeader>
                      <DialogTitle>Board umbenennen</DialogTitle>
                    </DialogHeader>
                    <div className="grid gap-4 py-4">
                      <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="name" className="text-right">
                          Name
                        </Label>
                        <Input
                          id="name"
                          defaultValue={x.Name}
                          onChange={(e) => setName(e.target.value)}
                          className="col-span-3"
                        />
                      </div>
                    </div>
                    <DialogFooter>
                      <Button
                        type="submit"
                        onClick={async () => await handleRename(x.ID)}
                      >
                        Speichern
                      </Button>
                    </DialogFooter>
                  </DialogContent>
                </Dialog>
                <AlertDialog>
                  <AlertDialogTrigger asChild>
                    <Button variant="destructive">Board Löschen</Button>
                  </AlertDialogTrigger>
                  <AlertDialogContent>
                    <AlertDialogHeader>
                      <AlertDialogTitle>
                        Bist du wirklich sicher?
                      </AlertDialogTitle>
                      <AlertDialogDescription>
                        Dieser Vorgang kann nicht Rückgängig gemacht werden!
                      </AlertDialogDescription>
                    </AlertDialogHeader>
                    <AlertDialogFooter>
                      <AlertDialogCancel>Abbrechen</AlertDialogCancel>
                      <AlertDialogAction
                        onClick={async () => handleDelete(x.ID)}
                      >
                        Löschen
                      </AlertDialogAction>
                    </AlertDialogFooter>
                  </AlertDialogContent>
                </AlertDialog>
              </div>
            </CardFooter>
          </Card>
        ))}
      </div>
    </>
  );
}
