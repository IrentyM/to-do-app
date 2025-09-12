export namespace dto {
	
	export class Task {
	    id: number;
	    title: string;
	    done: boolean;
	    created_at: string;
	    deadline?: string;
	    priority: number;
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.done = source["done"];
	        this.created_at = source["created_at"];
	        this.deadline = source["deadline"];
	        this.priority = source["priority"];
	    }
	}

}

export namespace model {
	
	export class Task {
	    ID: number;
	    Title: string;
	    Done: boolean;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    Deadline?: any;
	    Priority: number;
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Title = source["Title"];
	        this.Done = source["Done"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.Deadline = this.convertValues(source["Deadline"], null);
	        this.Priority = source["Priority"];
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

