import BackButton from "@/components/BackButton";
import { Button } from "@/components/ui/button";
import { Mitarbeiter } from "@api/db";
import type { db } from "@wails/go/models";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router";
import MitarbeiterForm from "./_components/mitarbeiter-form";

export default function MitarbeiterBearbeiten() {
  const { id } = useParams();
  const [mitarbeiter, setMitarbeiter] = useState<db.Mitarbeiter | undefined>(
    undefined
  );
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    async function x() {
      if (id == null) return;
      setLoading(true);
      const i = parseInt(id);
      const res = await Mitarbeiter.Get(i);
      setMitarbeiter(res);
      setLoading(false);
    }

    x();
  }, [id]);

  const handleDelete = async () => {
    if (id == null) return;
    const i = parseInt(id);
    await Mitarbeiter.Delete(i);
    navigate("/Mitarbeiter");
  };

  return (
    <>
      <BackButton href={"/Mitarbeiter/" + id} />
      <h1 className="text-center">Bearbeiten</h1>
      {!loading && mitarbeiter && <MitarbeiterForm mitarbeiter={mitarbeiter} />}
      <Button variant={"destructive"} onClick={handleDelete}>
        Mitarbeiter löschen
      </Button>
    </>
  );
}
