import React from "react";

export interface AlertProps {
	className?: string;
	message: string;
	type: keyof typeof colors;
}

const colors = {
	success: "text-accent-text bg-green-400",
	failure: "text-accent-text bg-red-400",
};

export const Alert: React.FC<AlertProps> = ({ className, type, message }) => {
	let size = "md:w-9/12 sm:w-full";
	let pos = "z-10 inset-x-0 mx-auto mt-3";

	let css = `${className} ${size} ${pos} ${colors[type]}`;

	return (
		<div className={`flex justify-center py-2 px-6 rounded-lg text-sm font-bold ${css}`}>
			<span>{message}</span>
		</div>
	);
};
