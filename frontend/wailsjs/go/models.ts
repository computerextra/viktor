export namespace archive {
	
	export class ArchiveResult {
	    Id: number;
	    Title: string;
	    Body: string;
	
	    static createFrom(source: any = {}) {
	        return new ArchiveResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.Title = source["Title"];
	        this.Body = source["Body"];
	    }
	}

}

export namespace db {
	
	export class Ansprechpartner {
	    ID: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	    // Go type: gorm
	    DeletedAt: any;
	    Name: string;
	    Telefon?: string;
	    Mobil?: string;
	    Mail?: string;
	    LieferantenId: number;
	
	    static createFrom(source: any = {}) {
	        return new Ansprechpartner(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.Name = source["Name"];
	        this.Telefon = source["Telefon"];
	        this.Mobil = source["Mobil"];
	        this.Mail = source["Mail"];
	        this.LieferantenId = source["LieferantenId"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Mitarbeiter {
	    ID: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	    // Go type: gorm
	    DeletedAt: any;
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
	    Geburtstag: sql.NullTime;
	    Paypal: boolean;
	    Abonniert: boolean;
	    Geld?: string;
	    Pfand?: string;
	    Dinge?: string;
	    Abgeschickt: sql.NullTime;
	    Bild1?: string;
	    Bild2?: string;
	    Bild3?: string;
	    Bild1Date: sql.NullTime;
	    Bild2Date: sql.NullTime;
	    Bild3Date: sql.NullTime;
	
	    static createFrom(source: any = {}) {
	        return new Mitarbeiter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.Name = source["Name"];
	        this.Short = source["Short"];
	        this.Gruppenwahl = source["Gruppenwahl"];
	        this.InternTelefon1 = source["InternTelefon1"];
	        this.InternTelefon2 = source["InternTelefon2"];
	        this.FestnetzPrivat = source["FestnetzPrivat"];
	        this.FestnetzBusiness = source["FestnetzBusiness"];
	        this.HomeOffice = source["HomeOffice"];
	        this.MobilBusiness = source["MobilBusiness"];
	        this.MobilPrivat = source["MobilPrivat"];
	        this.Email = source["Email"];
	        this.Azubi = source["Azubi"];
	        this.Geburtstag = this.convertValues(source["Geburtstag"], sql.NullTime);
	        this.Paypal = source["Paypal"];
	        this.Abonniert = source["Abonniert"];
	        this.Geld = source["Geld"];
	        this.Pfand = source["Pfand"];
	        this.Dinge = source["Dinge"];
	        this.Abgeschickt = this.convertValues(source["Abgeschickt"], sql.NullTime);
	        this.Bild1 = source["Bild1"];
	        this.Bild2 = source["Bild2"];
	        this.Bild3 = source["Bild3"];
	        this.Bild1Date = this.convertValues(source["Bild1Date"], sql.NullTime);
	        this.Bild2Date = this.convertValues(source["Bild2Date"], sql.NullTime);
	        this.Bild3Date = this.convertValues(source["Bild3Date"], sql.NullTime);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Geburtstagsliste {
	    Vergangen: Mitarbeiter[];
	    Heute: Mitarbeiter[];
	    Zukunft: Mitarbeiter[];
	
	    static createFrom(source: any = {}) {
	        return new Geburtstagsliste(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Vergangen = this.convertValues(source["Vergangen"], Mitarbeiter);
	        this.Heute = this.convertValues(source["Heute"], Mitarbeiter);
	        this.Zukunft = this.convertValues(source["Zukunft"], Mitarbeiter);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Post {
	    ID: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	    // Go type: gorm
	    DeletedAt: any;
	    KanbanId: number;
	    Name: string;
	    Description?: string;
	    Status: string;
	    Importance: string;
	
	    static createFrom(source: any = {}) {
	        return new Post(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.KanbanId = source["KanbanId"];
	        this.Name = source["Name"];
	        this.Description = source["Description"];
	        this.Status = source["Status"];
	        this.Importance = source["Importance"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Kanban {
	    ID: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	    // Go type: gorm
	    DeletedAt: any;
	    Name: string;
	    UserId: number;
	    Posts: Post[];
	
	    static createFrom(source: any = {}) {
	        return new Kanban(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.Name = source["Name"];
	        this.UserId = source["UserId"];
	        this.Posts = this.convertValues(source["Posts"], Post);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Lieferant {
	    ID: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	    // Go type: gorm
	    DeletedAt: any;
	    Firma: string;
	    Kundennummer?: string;
	    Webseite?: string;
	    Ansprechpartner: Ansprechpartner[];
	
	    static createFrom(source: any = {}) {
	        return new Lieferant(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.Firma = source["Firma"];
	        this.Kundennummer = source["Kundennummer"];
	        this.Webseite = source["Webseite"];
	        this.Ansprechpartner = this.convertValues(source["Ansprechpartner"], Ansprechpartner);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	export class User {
	    ID: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	    // Go type: gorm
	    DeletedAt: any;
	    Password: string;
	    Mail: string;
	    MitarbeiterId: number;
	    Mitarbeiter: Mitarbeiter;
	    Boards: Kanban[];
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.Password = source["Password"];
	        this.Mail = source["Mail"];
	        this.MitarbeiterId = source["MitarbeiterId"];
	        this.Mitarbeiter = this.convertValues(source["Mitarbeiter"], Mitarbeiter);
	        this.Boards = this.convertValues(source["Boards"], Kanban);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace sagedb {
	
	export class SearchResult {
	    SG_Adressen_PK: number;
	    Suchbegriff?: string;
	    KundNr?: string;
	    LiefNr?: string;
	    Homepage?: string;
	    Telefon1?: string;
	    Telefon2?: string;
	    Mobiltelefon1?: string;
	    Mobiltelefon2?: string;
	    EMail1?: string;
	    EMail2?: string;
	    KundUmsatz?: number;
	    LiefUmsatz?: number;
	
	    static createFrom(source: any = {}) {
	        return new SearchResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.SG_Adressen_PK = source["SG_Adressen_PK"];
	        this.Suchbegriff = source["Suchbegriff"];
	        this.KundNr = source["KundNr"];
	        this.LiefNr = source["LiefNr"];
	        this.Homepage = source["Homepage"];
	        this.Telefon1 = source["Telefon1"];
	        this.Telefon2 = source["Telefon2"];
	        this.Mobiltelefon1 = source["Mobiltelefon1"];
	        this.Mobiltelefon2 = source["Mobiltelefon2"];
	        this.EMail1 = source["EMail1"];
	        this.EMail2 = source["EMail2"];
	        this.KundUmsatz = source["KundUmsatz"];
	        this.LiefUmsatz = source["LiefUmsatz"];
	    }
	}
	export class User {
	    Name?: string;
	    Vorname?: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Vorname = source["Vorname"];
	    }
	}

}

export namespace sql {
	
	export class NullTime {
	    // Go type: time
	    Time: any;
	    Valid: boolean;
	
	    static createFrom(source: any = {}) {
	        return new NullTime(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Time = this.convertValues(source["Time"], null);
	        this.Valid = source["Valid"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace userdata {
	
	export class UserData {
	    Name?: string;
	    Mail?: string;
	    Id?: number;
	    UserId?: number;
	
	    static createFrom(source: any = {}) {
	        return new UserData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Mail = source["Mail"];
	        this.Id = source["Id"];
	        this.UserId = source["UserId"];
	    }
	}

}

