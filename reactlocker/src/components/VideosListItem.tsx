import React from "react";
import { Color, Video } from "../shared/types";
import { ColorBadge } from "./ColorBadge";
import { RightArrow } from "./Svg";

export interface VideoListItemProps {
	className?: string;
	video: Video;
	url: string;
	color: Color;
}

export const VideoListItem: React.FC<VideoListItemProps> = ({ video, className, url, color }) => {
	const css = `${className} hover:bg-primary-600 rounded-md flex justify-between cursor-pointer overflow-hidden`;
	const imageSize = "md:h-20 md:w-32 h-16 w-32";
	const textSize = "sm:text-md text-md";

	const open = () => {
		window.open(url, "_blank");
	};

	return (
		<div className={css} onClick={open}>
			<div className="flex p-1 overflow-hidden">
				<img
					src={video.thumbnailUrl}
					alt="Thumbnail"
					className={`${imageSize} flex-shrink-0 rounded-lg object-cover`}
				/>
				<div className="pl-3 flex flex-col">
					<span className={`${textSize} font-semibold whitespace-nowrap`}>{video.title}</span>
					<span className="text-sm text-primary-text-200">{video.created.toDateString()}</span>
					<div>
						<ColorBadge className="mt-2" color={color}></ColorBadge>
					</div>
				</div>
			</div>
			<div className="mr-2 my-auto select-none ml-4">
				<RightArrow size={24}></RightArrow>
			</div>
		</div>
	);
};
