export namespace main {
	
	export class Mitarbeiter {
	    id?: string;
	    name?: string;
	    short?: string;
	    image?: boolean;
	    sex?: string;
	    focus?: string;
	    mail?: string;
	    abteilungId?: string;
	    einkaufId?: string;
	    Azubi?: boolean;
	    // Go type: time
	    Geburtstag?: any;
	    Gruppenwahl?: string;
	    HomeOffice?: string;
	    Mobil_Business?: string;
	    Mobil_Privat?: string;
	    Telefon_Business?: string;
	    Telefon_Intern_1?: string;
	    Telefon_Intern_2?: string;
	    Telefon_Privat?: string;
	    // Go type: ent
	    edges: any;
	    Diff: number;
	
	    static createFrom(source: any = {}) {
	        return new Mitarbeiter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.short = source["short"];
	        this.image = source["image"];
	        this.sex = source["sex"];
	        this.focus = source["focus"];
	        this.mail = source["mail"];
	        this.abteilungId = source["abteilungId"];
	        this.einkaufId = source["einkaufId"];
	        this.Azubi = source["Azubi"];
	        this.Geburtstag = this.convertValues(source["Geburtstag"], null);
	        this.Gruppenwahl = source["Gruppenwahl"];
	        this.HomeOffice = source["HomeOffice"];
	        this.Mobil_Business = source["Mobil_Business"];
	        this.Mobil_Privat = source["Mobil_Privat"];
	        this.Telefon_Business = source["Telefon_Business"];
	        this.Telefon_Intern_1 = source["Telefon_Intern_1"];
	        this.Telefon_Intern_2 = source["Telefon_Intern_2"];
	        this.Telefon_Privat = source["Telefon_Privat"];
	        this.edges = this.convertValues(source["edges"], null);
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
	export class Geburtstag {
	    Heute: Mitarbeiter[];
	    Zukunft: Mitarbeiter[];
	    Vergangenheit: Mitarbeiter[];
	
	    static createFrom(source: any = {}) {
	        return new Geburtstag(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Heute = this.convertValues(source["Heute"], Mitarbeiter);
	        this.Zukunft = this.convertValues(source["Zukunft"], Mitarbeiter);
	        this.Vergangenheit = this.convertValues(source["Vergangenheit"], Mitarbeiter);
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

