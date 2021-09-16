import React from "react";

export interface CheckboxProps {
	className?: string;
	message: string;
	checked: boolean;
	setChecked: (checked: boolean) => void;
}

export const Checkbox: React.FC<CheckboxProps> = ({ className, message, checked, setChecked }) => {
	return (
		<label className={`${className} flex items-center`}>
			<input
				onClick={() => setChecked(!checked)}
				type="checkbox"
				value={(checked ? 1 : 0).toString()}
				className="form-tick rounded h-5 w-5 border border-gray-300 rounded-md checked:bg-accent focus:outline-none"
			/>
			<span className="ml-2 text-sm">{message}</span>
		</label>
	);
};
