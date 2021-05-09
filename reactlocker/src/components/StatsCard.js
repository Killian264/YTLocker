import React from 'react';
import PropTypes from 'prop-types';

const border = "border-t-8 border-accent border-solid";

export const StatsCard = ({ children, header, count, measurement, classes="", ...props }) => {
	return (
		<div className={`bg-primary-600 m-1 px-4 pb-2 pt-3 rounded-md ${border} ${classes}`} {...props} >
			<span className="text-lg">{header}</span>
			<div>
				<span className="text-xl text-accent">{count}</span>
				<span className="text-lg font-bold tracking-wider"> {measurement}</span>
			</div>
		</div>
	);
};

StatsCard.propTypes = {
	header: PropTypes.string,
	count: PropTypes.string,
	measurement: PropTypes.string,
	classes: PropTypes.string,
};
