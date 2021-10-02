import React from "react";

export interface InputProps extends React.HTMLAttributes<HTMLInputElement> {
	className?: string;
	value: string;
	type?: string;
	disabled?: boolean;
}

export const Input: React.FC<InputProps> = ({ className, value, type = "", disabled = false, ...props }) => {
	const bg = `bg-primary-600`;
	const ring = `focus:outline-none focus:ring-2 focus:ring-accent`;
	const text = `text-primary-text-200 focus:text-primary-text`;
	const c = `${className} ${ring} ${bg} ${text} w-full py-2 px-4 rounded-md placeholder-primary-text-300`;

	return <input className={c} type={type} value={value} disabled={disabled} {...props} />;
};
