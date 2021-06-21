import { useState } from "react";
import { Button } from "../components/Button";
import { Card } from "../components/Card";
import { Input } from "../components/Input";
import { LeftArrow } from "../components/Svg";
import { TextBox } from "./TextBox";

export interface ChannelSubscribeCardProps {
	className?: string;
	CreateClick: (title: string, description: string) => void;
	BackClick: () => void;
}

export const PlaylistCreateCard: React.FC<ChannelSubscribeCardProps> = ({
	className,
	CreateClick,
	BackClick,
}) => {
	const [title, setTitle] = useState("");
	const [description, setDescription] = useState("");

	return (
		<Card className={`${className} flex flex-col justify-content-between`}>
			<div className="flex justify-between -mb-1 -mt-1">
				<span className="leading-none text-2xl font-semibold">Create Playlist</span>
			</div>
			<a target="_blank" rel="noreferrer">
				<div className={`col-span-6 rounded-lg object-cover w-full bg-black h-40`} />
			</a>
			<div className="mt-3">
				<Input
					placeholder="Title"
					value={title}
					className="mb-3"
					onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
						setTitle(e.target.value);
					}}
				></Input>
				<TextBox
					placeholder="Description"
					value={description}
					className="mb-3"
					onChange={(e: React.ChangeEvent<HTMLTextAreaElement>) => {
						setDescription(e.target.value);
					}}
				></TextBox>
				<div className="pt-4 flex justify-between">
					<Button size="medium" color="secondary" onClick={BackClick}>
						Back
					</Button>
					<Button
						size="medium"
						color="primary"
						disabled={title === ""}
						onClick={() => {
							CreateClick(title, description);
							BackClick();
						}}
					>
						Create
					</Button>
				</div>
			</div>
		</Card>
	);
};
