export namespace db {
	
	export class Notebook {
	    id: number;
	    Title: string;
	
	    static createFrom(source: any = {}) {
	        return new Notebook(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.Title = source["Title"];
	    }
	}

}

