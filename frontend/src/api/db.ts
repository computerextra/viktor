import {
  ChangePassword,
  CreateAnsprechpartner,
  CreateLieferant,
  CreateMitarbeiter,
  CreateUser,
  DeleteAnsprechpartner,
  DeleteEinkauf,
  DeleteLieferant,
  DeleteMitarbeiter,
  DeleteUser,
  GetAllMitarbeiter,
  GetAllMitarbeiterMitEinkauf,
  GetAnsprechpartner,
  GetAnsprechpartnerFromLieferant,
  GetEinkaufsliste,
  GetGeburtstagsliste,
  GetLieferant,
  GetLieferanten,
  GetMitarbeiter,
  GetMitarbeiterMitEinkauf,
  GetUser,
  SkipEinkauf,
  UpdateAnsprechpartner,
  UpdateEinkauf,
  UpdateLieferant,
  UpdateMitarbeiter,
} from "@bindings/viktor/backend/app";
import { z } from "zod";

// Ansprechpartner
export const AnsprechpartnerParams = z.object({
  Name: z.string(),
  Telefon: z.string().optional(),
  Mobil: z.string().optional(),
  Mail: z.string().optional(),
  LieferantenId: z.string(),
});
export type AnsprechpartnerParams = z.infer<typeof AnsprechpartnerParams>;

export class Ansprechpartner {
  static async Create(params: AnsprechpartnerParams) {
    return await CreateAnsprechpartner(
      params.Name,
      params.Telefon ? params.Telefon : null,
      params.Mobil ? params.Mobil : null,
      params.Mail ? params.Mail : null,
      params.LieferantenId
    );
  }

  static async Update(id: string, params: AnsprechpartnerParams) {
    return await UpdateAnsprechpartner(
      id,
      params.Name,
      params.Telefon ? params.Telefon : null,
      params.Mobil ? params.Mobil : null,
      params.Mail ? params.Mail : null
    );
  }

  static async Get(id: string) {
    return await GetAnsprechpartner(id);
  }

  static async GetAll(lieferantenId: string) {
    return await GetAnsprechpartnerFromLieferant(lieferantenId);
  }

  static async Delete(id: string) {
    return await DeleteAnsprechpartner(id);
  }
}

export const LiefertantenParams = z.object({
  Firma: z.string(),
  Kundennummer: z.string().optional(),
  Webseite: z.string().optional(),
});
export type LiefertantenParams = z.infer<typeof LiefertantenParams>;

export class Lieferant {
  static async Create(params: LiefertantenParams) {
    return await CreateLieferant(
      params.Firma,
      params.Kundennummer ? params.Kundennummer : null,
      params.Webseite ? params.Webseite : null
    );
  }

  static async Update(id: string, params: LiefertantenParams) {
    return await UpdateLieferant(
      id,
      params.Firma,
      params.Kundennummer ? params.Kundennummer : null,
      params.Webseite ? params.Webseite : null
    );
  }

  static async Get(id: string) {
    return await GetLieferant(id);
  }

  static async GetAll() {
    return await GetLieferanten();
  }

  static async Delete(id: string) {
    return await DeleteLieferant(id);
  }
}

type CreateMitarbeiterParams = {
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
  Azubi: boolean;
  Geburtstag?: Date;
};

export class Mitarbeiter {
  static async Create(params: CreateMitarbeiterParams) {
    return await CreateMitarbeiter(
      params.Name,
      params.Short ? params.Short : null,
      params.Gruppenwahl ? params.Gruppenwahl : null,
      params.InternTelefon1 ? params.InternTelefon1 : null,
      params.InternTelefon2 ? params.InternTelefon2 : null,
      params.FestnetzPrivat ? params.FestnetzPrivat : null,
      params.FestnetzBusiness ? params.FestnetzBusiness : null,
      params.HomeOffice ? params.HomeOffice : null,
      params.MobilBusiness ? params.MobilBusiness : null,
      params.MobilPrivat ? params.MobilPrivat : null,
      params.Email ? params.Email : null,
      params.Azubi,
      params.Geburtstag
        ? params.Geburtstag.toLocaleString("de-DE", {
            day: "2-digit",
            month: "2-digit",
            year: "numeric",
          })
        : ""
    );
  }

  static async Update(id: string, params: CreateMitarbeiterParams) {
    return await UpdateMitarbeiter(
      id,
      params.Name,
      params.Short ? params.Short : null,
      params.Gruppenwahl ? params.Gruppenwahl : null,
      params.InternTelefon1 ? params.InternTelefon1 : null,
      params.InternTelefon2 ? params.InternTelefon2 : null,
      params.FestnetzPrivat ? params.FestnetzPrivat : null,
      params.FestnetzBusiness ? params.FestnetzBusiness : null,
      params.HomeOffice ? params.HomeOffice : null,
      params.MobilBusiness ? params.MobilBusiness : null,
      params.MobilPrivat ? params.MobilPrivat : null,
      params.Email ? params.Email : null,
      params.Azubi,
      params.Geburtstag
        ? params.Geburtstag.toLocaleString("de-DE", {
            day: "2-digit",
            month: "2-digit",
            year: "numeric",
          })
        : ""
    );
  }
  static async UpdateEinkauf(
    id: string,
    params: {
      Paypal: boolean;
      Abonniert: boolean;
      Geld?: string;
      Pfand?: string;
      Dinge?: string;
      Bild1: boolean;
      Bild2: boolean;
      Bild3: boolean;
    }
  ) {
    return await UpdateEinkauf(
      id,
      params.Paypal,
      params.Abonniert,
      params.Geld ? params.Geld : null,
      params.Pfand ? params.Pfand : null,
      params.Dinge ? params.Dinge : null,
      params.Bild1,
      params.Bild2,
      params.Bild3
    );
  }

  static async EinkaufSkip(id: string) {
    return await SkipEinkauf(id);
  }

  static async EinkaufDelete(id: string) {
    return await DeleteEinkauf(id);
  }

  static async Get(id: string) {
    return await GetMitarbeiter(id);
  }

  static async GetAll() {
    return await GetAllMitarbeiter();
  }

  static async GetAllWithEinkauf() {
    return await GetAllMitarbeiterMitEinkauf();
  }

  static async GetWithEinkauf(id: string) {
    return await GetMitarbeiterMitEinkauf(id);
  }

  static async Einkauf() {
    return await GetEinkaufsliste();
  }

  static async Geburtstag() {
    return await GetGeburtstagsliste();
  }

  static async Delete(id: string) {
    return await DeleteMitarbeiter(id);
  }
}

export class User {
  static async Create(Mail: string, Password: string): Promise<boolean> {
    return await CreateUser(Mail, Password);
  }

  static async Get(id: string) {
    return await GetUser(id);
  }

  static async ChangePassword(id: string, New: string, Old: string) {
    return await ChangePassword(id, New, Old);
  }

  static async Delete(id: string) {
    return await DeleteUser(id);
  }
}
