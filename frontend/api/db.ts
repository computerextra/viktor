import {
  ChangePassword,
  CheckUser,
  CreateAnsprechpartner,
  CreateLieferant,
  CreateMitarbeiter,
  CreateUser,
  DeleteAnsprechpartner,
  DeleteLieferant,
  DeleteMitarbeiter,
  DeleteUser,
  GetAllAnsprechpartner,
  GetAllMitarbeiter,
  GetAllMitarbeiterEinkauf,
  GetAnsprechpartner,
  GetEinkaufsliste,
  GetGeburtstagsliste,
  GetLieferant,
  GetLieferanten,
  GetMitarbeiter,
  GetUser,
  UpdateAnsprechpartner,
  UpdateEinkauf,
  UpdateLieferant,
  UpdateMitarbeiter,
} from "@wails/go/main/App";

// Ansprechpartner

type AnsprechpartnerParams = {
  Name: string;
  Telefon?: string;
  Mobil?: string;
  Mail?: string;
  LieferantenId: number;
};

export class Ansprechpartner {
  static async Create(params: AnsprechpartnerParams) {
    return await CreateAnsprechpartner(
      params.Name,
      params.Telefon,
      params.Mobil,
      params.Mail,
      params.LieferantenId
    );
  }

  static async Update(id: number, params: AnsprechpartnerParams) {
    return await UpdateAnsprechpartner(
      id,
      params.Name,
      params.Telefon,
      params.Mobil,
      params.Mail
    );
  }

  static async Get(id: number) {
    return await GetAnsprechpartner(id);
  }

  static async GetAll() {
    return await GetAllAnsprechpartner();
  }

  static async Delete(id: number) {
    return await DeleteAnsprechpartner(id);
  }
}

type LiefertantenParams = {
  Firma: string;
  Kundennummer?: string;
  Webseite?: string;
};

export class Lieferant {
  static async Create(params: LiefertantenParams) {
    return await CreateLieferant(
      params.Firma,
      params.Kundennummer,
      params.Webseite
    );
  }

  static async Update(id: number, params: LiefertantenParams) {
    return await UpdateLieferant(
      id,
      params.Firma,
      params.Kundennummer,
      params.Webseite
    );
  }

  static async Get(id: number) {
    return await GetLieferant(id);
  }

  static async GetAll() {
    return await GetLieferanten();
  }

  static async Delete(id: number) {
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
      params.Short,
      params.Gruppenwahl,
      params.InternTelefon1,
      params.InternTelefon2,
      params.FestnetzPrivat,
      params.FestnetzBusiness,
      params.HomeOffice,
      params.MobilBusiness,
      params.MobilPrivat,
      params.Email,
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

  static async Update(id: number, params: CreateMitarbeiterParams) {
    return await UpdateMitarbeiter(
      id,
      params.Name,
      params.Short,
      params.Gruppenwahl,
      params.InternTelefon1,
      params.InternTelefon2,
      params.FestnetzPrivat,
      params.FestnetzBusiness,
      params.HomeOffice,
      params.MobilBusiness,
      params.MobilPrivat,
      params.Email,
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
    id: number,
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
      params.Geld,
      params.Pfand,
      params.Dinge,
      params.Bild1,
      params.Bild2,
      params.Bild3
    );
  }

  static async Get(id: number) {
    return await GetMitarbeiter(id);
  }

  static async GetAll() {
    return await GetAllMitarbeiter();
  }

  static async GetAllEinkauf() {
    return await GetAllMitarbeiterEinkauf();
  }

  static async Einkauf() {
    return await GetEinkaufsliste();
  }

  static async Geburtstag() {
    return await GetGeburtstagsliste();
  }

  static async Delete(id: number) {
    return await DeleteMitarbeiter(id);
  }
}

export class User {
  static async Create(Mail: string, Password: string): Promise<string> {
    return await CreateUser(Mail, Password);
  }

  static async Get(id: number) {
    return await GetUser(id);
  }

  static async Check(Mail: string, Password: string): Promise<boolean> {
    return await CheckUser(Mail, Password);
  }

  static async ChangePassword(id: number, New: string, Old: string) {
    return await ChangePassword(id, New, Old);
  }

  static async Delete(id: number) {
    return await DeleteUser(id);
  }
}
