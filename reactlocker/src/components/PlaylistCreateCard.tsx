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
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="flex items-center gap-2">
					<div onClick={BackClick} className="cursor-pointer -m-1">
						<LeftArrow size={32} strokeWidth={2}></LeftArrow>
					</div>
					<span className="leading-none text-2xl font-semibold">Create Playlist</span>
				</div>
			</div>
			<div>
				<div className="mb-1">
					<span className="text-2xl font-semibold leading-none">Title:</span>
				</div>
				<Input
					value={title}
					className="mb-3"
					onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
						setTitle(e.target.value);
					}}
				></Input>
				<div className="mb-1">
					<span className="text-2xl font-semibold leading-none">Description</span>
				</div>
				<TextBox
					value={description}
					className="mb-3"
					onChange={(e: React.ChangeEvent<HTMLTextAreaElement>) => {
						setDescription(e.target.value);
					}}
				></TextBox>
				<div className="mt-3 flex justify-between">
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
