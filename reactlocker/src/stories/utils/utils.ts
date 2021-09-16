// FROM https://github.com/benawad/dogehouse/

export const sRadio = (options: string[]) => ({
	control: {
		type: "inline-radio",
		options: options,
	},
	defaultValue: options[0],
});

export const sString = (defaultValue = "") => ({
	control: {
		type: "text",
	},
	defaultValue: defaultValue,
});

export const sBoolean = (defaultValue = false) => ({
	control: {
		type: "boolean",
	},
	defaultValue: defaultValue,
});
