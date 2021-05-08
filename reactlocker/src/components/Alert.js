import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';


const colors = {
	success: "text-accent-text bg-green-400",
	failure: "text-accent-text bg-red-400",
}


export const Alert = ({ className, children, type="success", ...props }) => {

	let size = "md:w-9/12 sm:w-full"
	let pos = "z-10 fixed inset-x-0 mx-auto mt-3"
	let css = `${size} ${pos} ${colors[type]} ${className}`

	return (
		<div
			type="button"
			className={`flex justify-center py-2 px-6 rounded-lg text-sm font-bold ${css}`}
			{...props}
		>
			<span>{children}</span>
		</div>
	);
};

Alert.propTypes = {
	type: PropTypes.oneOf(['success', 'failure']),
	className: PropTypes.string,
};
