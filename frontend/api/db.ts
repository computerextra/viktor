import type { db } from "@wails/go/models";
import {
  Create as APICreate,
  Update as APIUpdate,
  Read as APIRead,
  Einkaufsliste as APIEinkauf,
  Geburtstagsliste as APIGeburtstag,
  Delete as APIDelete,
} from "@wails/go/main/App";

type Models =
  | "Ansprechpartner"
  | "Lieferant"
  | "Mitarbeiter"
  | "User"
  | "Version";

interface AnsprechpartnerParams {
  Name: string;
  Telefon?: string;
  Mobil?: string;
  Mail?: string;
  LieferantenId?: string;
}

interface Ansprechparnter extends AnsprechpartnerParams {
  Id: string;
}

interface LieferantParams {
  Firma: string;
  Kundennummer?: string;
  Webseite?: string;
}

interface Lieferant extends LieferantParams {
  Id: string;
}

type MitarbeiterParams = {
  Name: string;
  Short?: string;
  Gruppenwahl?: string;
  InternTelefon1?: string;
  InternTelefon2?: string;
  FestnetzPrivat?: string;
  FestnetzBusiness?: string;
  HomeOffice?: string;
  MobilBusiness?: string;
  MobilPrivat?: string;
  Email?: string;
  Azubi?: boolean;
  Geburtstag?: Date;
  Paypal?: boolean;
  Abonniert?: boolean;
  Geld?: string;
  Pfand?: string;
  Dinge?: string;
  Abgeschickt?: string;
  Bild1?: string;
  Bild2?: string;
  Bild3?: string;
  Bild1Date?: Date;
  Bild2Date?: Date;
  Bild3Date?: Date;
};

interface UserParams {
  Password: string;
  Mail: string;
  Active: boolean;
}
interface User extends UserParams {
  Id: string;
  MitarbeiterId: string;
}

interface VersionParams {
  Current: number;
}

interface Version extends VersionParams {
  Id: number;
}

const Create = async (
  model: Models,
  params:
    | AnsprechpartnerParams
    | LieferantParams
    | MitarbeiterParams
    | UserParams
    | VersionParams
): Promise<boolean> => {
  const response = await APICreate(model, params);
  return response;
};

const Update = async (
  model: Models,
  params:
    | AnsprechpartnerParams
    | LieferantParams
    | MitarbeiterParams
    | UserParams
    | VersionParams,
  id: string | number
): Promise<boolean> => {
  const response = await APIUpdate(
    model,
    params,
    typeof id == "string" ? id : null,
    typeof id == "number" ? id : null
  );
  return response;
};

const Read = async (
  model: Models,
  id: string | number
): Promise<
  Ansprechparnter[] | Lieferant[] | db.MitarbeiterModel[] | User[] | Version[]
> => {
  const results = await APIRead(
    model,
    typeof id == "string" ? id : null,
    typeof id == "number" ? id : null
  );
  return results;
};

const Delete = async (model: Models, id: string | number): Promise<boolean> => {
  const response = await APIDelete(
    model,
    typeof id == "string" ? id : null,
    typeof id == "number" ? id : null
  );
  return response;
};

const Einkaufsliste = async (): Promise<Array<db.MitarbeiterModel>> => {
  const res = await APIEinkauf();
  return res;
};

const GeburtstagsListe = async (): Promise<db.GeburtstagsListe> => {
  const res = await APIGeburtstag();
  return res;
};

export default {
  Create,
  Update,
  Read,
  Delete,
  Einkaufsliste,
  GeburtstagsListe,
};
