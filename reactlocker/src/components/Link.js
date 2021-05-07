import React from 'react';


export const Link = ({ children, ...props }) => {
	return (
		<span className="text-sm underline text-secondary-text cursor-pointer select-none" {...props} >{children}</span>
	);
};