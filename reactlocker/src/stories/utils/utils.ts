// FROM https://github.com/benawad/dogehouse/

export const sRadio = (choices: string[]) => ({
	control: {
		type: "inline-radio",
		options: choices,
	},
	defaultValue: choices[0],
});

export const sString = () => ({
	control: {
		type: "text",
	},
	defaultValue: "",
});

export const sBoolean = () => ({
	control: {
		type: "boolean",
	},
	defaultValue: false,
});
