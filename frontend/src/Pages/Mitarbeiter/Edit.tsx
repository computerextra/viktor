import BackButton from "@/components/BackButton";
import { Button } from "@/components/ui/button";
import { Mitarbeiter as MitarbeiterAPI } from "@api/db";
import type { Mitarbeiter } from "@bindings/viktor/db/models";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router";
import MitarbeiterForm from "./_components/mitarbeiter-form";

export default function MitarbeiterBearbeiten() {
  const { id } = useParams();
  const [mitarbeiter, setMitarbeiter] = useState<Mitarbeiter | null>(null);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    async function x() {
      if (id == null) return;
      setLoading(true);
      const res = await MitarbeiterAPI.Get(id);
      setMitarbeiter(res);
      setLoading(false);
    }

    x();
  }, [id]);

  const handleDelete = async () => {
    if (id == null) return;
    await MitarbeiterAPI.Delete(id);
    navigate("/Mitarbeiter");
  };

  return (
    <>
      <BackButton href={"/Mitarbeiter/" + id} />
      <h1 className="text-center">{mitarbeiter?.Name} Bearbeiten</h1>
      {!loading && mitarbeiter && <MitarbeiterForm mitarbeiter={mitarbeiter} />}
      <Button variant={"destructive"} onClick={handleDelete}>
        Mitarbeiter löschen
      </Button>
    </>
  );
}
