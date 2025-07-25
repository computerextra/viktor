export namespace main {
	
	export class Mitarbeiter {
	    ID: string;
	    Name: string;
	    // Go type: sql
	    Short: any;
	    Image: boolean;
	    // Go type: sql
	    Sex: any;
	    // Go type: sql
	    Focus: any;
	    // Go type: sql
	    Mail: any;
	    // Go type: sql
	    Abteilungid: any;
	    // Go type: sql
	    Einkaufid: any;
	    Azubi: boolean;
	    // Go type: sql
	    Geburtstag: any;
	    // Go type: sql
	    Gruppenwahl: any;
	    // Go type: sql
	    Homeoffice: any;
	    // Go type: sql
	    MobilBusiness: any;
	    // Go type: sql
	    MobilPrivat: any;
	    // Go type: sql
	    TelefonBusiness: any;
	    // Go type: sql
	    TelefonIntern1: any;
	    // Go type: sql
	    TelefonIntern2: any;
	    // Go type: sql
	    TelefonPrivat: any;
	    Diff: number;
	
	    static createFrom(source: any = {}) {
	        return new Mitarbeiter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Short = this.convertValues(source["Short"], null);
	        this.Image = source["Image"];
	        this.Sex = this.convertValues(source["Sex"], null);
	        this.Focus = this.convertValues(source["Focus"], null);
	        this.Mail = this.convertValues(source["Mail"], null);
	        this.Abteilungid = this.convertValues(source["Abteilungid"], null);
	        this.Einkaufid = this.convertValues(source["Einkaufid"], null);
	        this.Azubi = source["Azubi"];
	        this.Geburtstag = this.convertValues(source["Geburtstag"], null);
	        this.Gruppenwahl = this.convertValues(source["Gruppenwahl"], null);
	        this.Homeoffice = this.convertValues(source["Homeoffice"], null);
	        this.MobilBusiness = this.convertValues(source["MobilBusiness"], null);
	        this.MobilPrivat = this.convertValues(source["MobilPrivat"], null);
	        this.TelefonBusiness = this.convertValues(source["TelefonBusiness"], null);
	        this.TelefonIntern1 = this.convertValues(source["TelefonIntern1"], null);
	        this.TelefonIntern2 = this.convertValues(source["TelefonIntern2"], null);
	        this.TelefonPrivat = this.convertValues(source["TelefonPrivat"], null);
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

