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
	
	export class Geburtstag {
	    Name: string;
	    // Go type: time
	    Geburtstag: any;
	    Diff: number;
	
	    static createFrom(source: any = {}) {
	        return new Geburtstag(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Geburtstag = this.convertValues(source["Geburtstag"], null);
	        this.Diff = source["Diff"];
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
	export class GeburtstagsListe {
	    Vergangen: Geburtstag[];
	    Heute: Geburtstag[];
	    Zukunft: Geburtstag[];
	
	    static createFrom(source: any = {}) {
	        return new GeburtstagsListe(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Vergangen = this.convertValues(source["Vergangen"], Geburtstag);
	        this.Heute = this.convertValues(source["Heute"], Geburtstag);
	        this.Zukunft = this.convertValues(source["Zukunft"], Geburtstag);
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
	export class MitarbeiterModel {
	    id: string;
	    Name: string;
	    Short?: string;
	    Gruppenwahl?: string;
	    Intern_telefon1?: string;
	    Intern_telefon2?: string;
	    Festnetz_privat?: string;
	    Festnetz_busines?: string;
	    Home_office?: string;
	    Mobil_buiness?: string;
	    Mobil_privat?: string;
	    Email?: string;
	    Azubi?: boolean;
	    // Go type: time
	    Geburtstag?: any;
	    Paypal?: boolean;
	    Abonniert?: boolean;
	    Geld?: string;
	    Pfand?: string;
	    Dinge?: string;
	    // Go type: time
	    Abgeschickt?: any;
	    Bild1?: string;
	    Bild2?: string;
	    Bild3?: string;
	    // Go type: time
	    Bild1Date?: any;
	    // Go type: time
	    Bild2Date?: any;
	    // Go type: time
	    Bild3Date?: any;
	
	    static createFrom(source: any = {}) {
	        return new MitarbeiterModel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.Name = source["Name"];
	        this.Short = source["Short"];
	        this.Gruppenwahl = source["Gruppenwahl"];
	        this.Intern_telefon1 = source["Intern_telefon1"];
	        this.Intern_telefon2 = source["Intern_telefon2"];
	        this.Festnetz_privat = source["Festnetz_privat"];
	        this.Festnetz_busines = source["Festnetz_busines"];
	        this.Home_office = source["Home_office"];
	        this.Mobil_buiness = source["Mobil_buiness"];
	        this.Mobil_privat = source["Mobil_privat"];
	        this.Email = source["Email"];
	        this.Azubi = source["Azubi"];
	        this.Geburtstag = this.convertValues(source["Geburtstag"], null);
	        this.Paypal = source["Paypal"];
	        this.Abonniert = source["Abonniert"];
	        this.Geld = source["Geld"];
	        this.Pfand = source["Pfand"];
	        this.Dinge = source["Dinge"];
	        this.Abgeschickt = this.convertValues(source["Abgeschickt"], null);
	        this.Bild1 = source["Bild1"];
	        this.Bild2 = source["Bild2"];
	        this.Bild3 = source["Bild3"];
	        this.Bild1Date = this.convertValues(source["Bild1Date"], null);
	        this.Bild2Date = this.convertValues(source["Bild2Date"], null);
	        this.Bild3Date = this.convertValues(source["Bild3Date"], null);
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

