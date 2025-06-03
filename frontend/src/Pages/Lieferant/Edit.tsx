import BackButton from "@/components/BackButton";
import { Button } from "@/components/ui/button";
import {
  Ansprechpartner as AnsprechpartnerAPI,
  Lieferant as LieferantAPI,
} from "@api/db";
import type { Ansprechpartner, Lieferant } from "bindings/viktor/db/models";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router";
import { AnsprechpartnerForm, LieferantenForm } from "./_components/forms";

export const EditLieferant = () => {
  const { id } = useParams();
  const [lieferant, setLieferant] = useState<Lieferant | null>(null);
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate();

  useEffect(() => {
    (async () => {
      if (id == null) return;
      setLoading(true);
      const res = await LieferantAPI.Get(id);
      setLieferant(res);
      setLoading(false);
    })();
  }, [id]);

  const handleDelete = async () => {
    if (id == null) return;
    setLoading(true);
    await LieferantAPI.Delete(id);
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

  const [ansprechpartner, setAnsprechpartner] =
    useState<Ansprechpartner | null>(null);
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate();

  useEffect(() => {
    (async () => {
      if (ap == null) return;
      setLoading(true);
      const res = await AnsprechpartnerAPI.Get(ap);
      setAnsprechpartner(res);
      setLoading(false);
    })();
  }, [ap]);

  const handleDelete = async () => {
    if (ap == null) return;
    setLoading(true);
    await AnsprechpartnerAPI.Delete(ap);
    navigate("/Lieferant/" + id);
    setLoading(false);
  };

  return (
    <>
      <BackButton href={`/Lieferant/${id}/${ap}`} />
      <h1 className="text-center">{ansprechpartner?.Name} Bearbeiten</h1>
      {!loading && id && ansprechpartner && (
        <AnsprechpartnerForm
          LieferantenId={id}
          ansprechpartner={ansprechpartner}
        />
      )}
      <Button className="mt-5" variant={"destructive"} onClick={handleDelete}>
        Ansprechpartner löschen
      </Button>
    </>
  );
};
