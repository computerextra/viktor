import BackButton from "@/components/BackButton";
import { Button } from "@/components/ui/button";
import { Ansprechpartner, Lieferant } from "@api/db";
import type { db } from "@wails/go/models";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router";
import { AnsprechpartnerForm, LieferantenForm } from "./_components/forms";

export const EditLieferant = () => {
  const { id } = useParams();
  const [lieferant, setLieferant] = useState<db.Lieferant | undefined>(
    undefined
  );
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate();

  useEffect(() => {
    (async () => {
      if (id == null) return;
      setLoading(true);
      const i = parseInt(id);
      const res = await Lieferant.Get(i);
      setLieferant(res);
      setLoading(false);
    })();
  }, [id]);

  const handleDelete = async () => {
    if (id == null) return;
    setLoading(true);
    const i = parseInt(id);
    await Lieferant.Delete(i);
    navigate("/Lieferant");
    setLoading(false);
  };

  return (
    <>
      <BackButton href={`/Lieferant/${id}`} />
      <h1 className="text-center">{lieferant?.Firma} Bearbeiten</h1>
      {!loading && lieferant && <LieferantenForm lieferant={lieferant} />}
      <Button
        className="mt-5"
        disabled={loading}
        variant={"destructive"}
        onClick={handleDelete}
      >
        Lieferant löschen
      </Button>
    </>
  );
};

export const EditAnsprechpartner = () => {
  const { id, ap } = useParams();

  const [ansprechpartner, setAnsprechpartner] = useState<
    db.Ansprechpartner | undefined
  >(undefined);
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate();

  useEffect(() => {
    (async () => {
      if (ap == null) return;
      setLoading(true);
      const i = parseInt(ap);
      const res = await Ansprechpartner.Get(i);
      setAnsprechpartner(res);
      setLoading(false);
    })();
  }, [ap]);

  const handleDelete = async () => {
    if (ap == null) return;
    setLoading(true);
    const i = parseInt(ap);
    await Ansprechpartner.Delete(i);
    navigate("/Lieferant/" + id);
    setLoading(false);
  };

  return (
    <>
      <BackButton href={`/Lieferant/${id}/${ap}`} />
      <h1 className="text-center">{ansprechpartner?.Name} Bearbeiten</h1>
      {!loading && id && ansprechpartner && (
        <AnsprechpartnerForm
          LieferantenId={parseInt(id)}
          ansprechpartner={ansprechpartner}
        />
      )}
      <Button className="mt-5" variant={"destructive"} onClick={handleDelete}>
        Ansprechpartner löschen
      </Button>
    </>
  );
};
