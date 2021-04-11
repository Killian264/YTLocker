// FROM https://github.com/benawad/dogehouse/

export const toEnum = (arr) => ({
	control: {
		type: "inline-radio",
		options: arr,
	},
});

export const toStr = () => ({
	control: {
	  type: "text",
	},
});

export const toBoolean = () => ({
	control: {
	  type: "boolean",
	},
});