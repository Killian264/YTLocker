import React from "react";

export interface TextBoxProps extends React.TextareaHTMLAttributes<HTMLTextAreaElement> {
	className?: string;
	value: string;
}

export const TextBox: React.FC<TextBoxProps> = ({ className, value, ...props }) => {
	const bg = `bg-primary-600`;
	const ring = `focus:outline-none focus:ring-2 focus:ring-accent`;
	const text = `text-primary-text-200 focus:text-primary-text`;
	const c = `${className} ${ring} ${bg} ${text} w-full py-2 px-4 rounded-md placeholder-primary-text-300 h-full`;

	return <textarea className={c} value={value} {...props} />;
};
