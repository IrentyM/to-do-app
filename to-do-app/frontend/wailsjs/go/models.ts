export namespace dto {
	
	export class Task {
	    id: number;
	    title: string;
	    done: boolean;
	    created_at: string;
	    deadline?: string;
	    priority: string;
	
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

